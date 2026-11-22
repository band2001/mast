package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Racer struct {
	Bib         string
	Name        string
	School      string
	Finish      int
	Seed        int
	SeedSize    int
	SeedRank    int
	OverallRank int
	Varsity     int
	SeedPoints  int
	Top30Points int
	DualPoints  int
	Total       int
}

type Team struct {
	School string
	Finish int
	Rank   int
}

func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.NewReplacer(
		" ", "", "_", "", "-", "",
		".", "", "#", "",
	).Replace(s)
}

func column(headers []string, names ...string) int {
	for i, h := range headers {
		for _, name := range names {
			if normalize(h) == normalize(name) {
				return i
			}
		}
	}
	return -1
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = -1
	return r.ReadAll()
}

func getField(row []string, index int) string {
	if index < 0 || index >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[index])
}

// Handles finishes such as 1, 1st, or 12th.
// Non-numeric finishes such as DNF receive no placement.
func parseFinish(s string) int {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return 0
	}

	var digits strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits.WriteRune(r)
		} else if digits.Len() > 0 {
			break
		} else if !unicode.IsSpace(r) {
			return 0
		}
	}

	n, err := strconv.Atoi(digits.String())
	if err != nil || n <= 0 {
		return 0
	}
	return n
}

func readRunOrder(path string) ([]*Racer, error) {
	rows, err := readCSV(path)
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, fmt.Errorf("run order CSV is empty")
	}

	headers := rows[0]
	bibCol := column(headers, "bib", "bib number", "bibnumber")
	nameCol := column(headers, "name", "racer name")
	schoolCol := column(headers, "school", "team")

	if bibCol < 0 || nameCol < 0 || schoolCol < 0 {
		return nil, fmt.Errorf("run order CSV must contain Bib, Name, and School columns")
	}

	var racers []*Racer
	seedCounts := make(map[string]int)
	seenBibs := make(map[string]bool)

	for _, row := range rows[1:] {
		bib := getField(row, bibCol)
		if bib == "" {
			continue
		}
		if seenBibs[bib] {
			return nil, fmt.Errorf("duplicate bib %s in run order", bib)
		}
		seenBibs[bib] = true

		school := getField(row, schoolCol)
		if school == "" {
			return nil, fmt.Errorf("bib %s has no school", bib)
		}

		// Each school's first racer is seed 1,
		// second racer is seed 2, and so on.
		seedCounts[normalize(school)]++

		r := &Racer{
			Bib:     bib,
			Name:    getField(row, nameCol),
			School:  school,
			Seed:    seedCounts[normalize(school)],
			Varsity: 0,
		}
		racers = append(racers, r)
	}

	return racers, nil
}

func readIndividualResults(path string, racers []*Racer) error {
	rows, err := readCSV(path)
	if err != nil {
		return err
	}
	if len(rows) < 2 {
		return fmt.Errorf("individual results CSV is empty")
	}

	headers := rows[0]
	bibCol := column(headers, "bib", "bib number", "bibnumber")
	nameCol := column(headers, "name", "racer name")
	schoolCol := column(headers, "school", "team")
	finishCol := column(headers, "finish", "place", "position")

	if bibCol < 0 || nameCol < 0 || schoolCol < 0 || finishCol < 0 {
		return fmt.Errorf("individual results CSV must contain Bib, Name, School, and Finish columns")
	}

	racerMap := make(map[string]*Racer)
	for _, r := range racers {
		racerMap[r.Bib] = r
	}

	seen := make(map[string]bool)

	for _, row := range rows[1:] {
		bib := getField(row, bibCol)
		if bib == "" {
			continue
		}

		r, exists := racerMap[bib]
		if !exists {
			return fmt.Errorf("bib %s appears in individual results but not in run order", bib)
		}
		if seen[bib] {
			return fmt.Errorf("duplicate bib %s in individual results", bib)
		}
		seen[bib] = true

		school := getField(row, schoolCol)
		if normalize(school) != normalize(r.School) {
			return fmt.Errorf("school mismatch for bib %s: run order=%s, results=%s",
				bib, r.School, school)
		}

		r.Name = getField(row, nameCol)
		r.Finish = parseFinish(getField(row, finishCol))
		r.Varsity = 0
	}

	// Every racer in the run order earns varsity selection points.
	for _, r := range racers {
		r.Varsity = 1
	}

	return nil
}

func readTeamResults(path string) ([]*Team, error) {
	rows, err := readCSV(path)
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, fmt.Errorf("team results CSV is empty")
	}

	headers := rows[0]
	schoolCol := column(headers, "school", "team")
	finishCol := column(headers, "finish", "place", "position")

	if schoolCol < 0 || finishCol < 0 {
		return nil, fmt.Errorf("team results CSV must contain School and Finish columns")
	}

	var teams []*Team
	seen := make(map[string]bool)

	for _, row := range rows[1:] {
		school := getField(row, schoolCol)
		if school == "" {
			continue
		}
		key := normalize(school)
		if seen[key] {
			return nil, fmt.Errorf("duplicate school %s in team results", school)
		}
		seen[key] = true

		teams = append(teams, &Team{
			School: school,
			Finish: parseFinish(getField(row, finishCol)),
		})
	}

	return teams, nil
}

// Assigns competition ranks.
// Tied finishes receive the same rank.
func assignRanks(racers []*Racer, getFinish func(*Racer) int,
	setRank func(*Racer, int)) {
	sorted := append([]*Racer(nil), racers...)

	sort.SliceStable(sorted, func(i, j int) bool {
		a, b := getFinish(sorted[i]), getFinish(sorted[j])
		if a == 0 {
			return false
		}
		if b == 0 {
			return true
		}
		return a < b
	})

	rank := 0
	lastFinish := -1

	for _, r := range sorted {
		finish := getFinish(r)
		if finish <= 0 {
			setRank(r, 0)
			continue
		}

		if finish != lastFinish {
			rank++
			lastFinish = finish
		}
		setRank(r, rank)
	}
}

func calculateInvitational(racers []*Racer) {
	seeds := make(map[int][]*Racer)

	for _, r := range racers {
		seeds[r.Seed] = append(seeds[r.Seed], r)
	}

	// Calculate placement points within each seed.
	for _, group := range seeds {
		for _, r := range group {
			r.SeedSize = len(group)
		}

		assignRanks(group,
			func(r *Racer) int { return r.Finish },
			func(r *Racer, rank int) { r.SeedRank = rank },
		)

		// Floor(seed size / 2) determines the scoring cutoff.
		cutoff := len(group) / 2

		for _, r := range group {
			if r.SeedRank > 0 && r.SeedRank <= cutoff {
				r.SeedPoints = cutoff - r.SeedRank + 1
			}
		}
	}

	// Calculate overall placement points.
	assignRanks(racers,
		func(r *Racer) int { return r.Finish },
		func(r *Racer, rank int) { r.OverallRank = rank },
	)

	// Top 30%, rounded down.
	cutoff := len(racers) * 30 / 100

	for _, r := range racers {
		if r.OverallRank > 0 && r.OverallRank <= cutoff {
			r.Top30Points = cutoff - r.OverallRank + 1
		}
	}
}

func calculateDual(racers []*Racer, teams []*Team) {
	// Rank teams by their team results.
	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].Finish == 0 {
			return false
		}
		if teams[j].Finish == 0 {
			return true
		}
		return teams[i].Finish < teams[j].Finish
	})

	rank := 0
	lastFinish := -1

	for _, t := range teams {
		if t.Finish <= 0 {
			continue
		}
		if t.Finish != lastFinish {
			rank++
			lastFinish = t.Finish
		}
		t.Rank = rank
	}

	teamRanks := make(map[string]int)
	for _, t := range teams {
		teamRanks[normalize(t.School)] = t.Rank
	}

	// Each skier receives points based on their school's
	// team placement: 3, 2, 1 for a three-team meet.
	numberOfTeams := len(teams)

	for _, r := range racers {
		teamRank := teamRanks[normalize(r.School)]
		if teamRank > 0 {
			r.DualPoints = numberOfTeams - teamRank + 1
		}
	}
}

func calculateTotals(racers []*Racer, raceType string) {
	for _, r := range racers {
		if raceType == "dual" {
			r.Total = r.Varsity + r.DualPoints
		} else {
			r.Total = r.Varsity + r.SeedPoints + r.Top30Points
		}
	}
}

func writeIndividualCSV(path string, racers []*Racer, raceType string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	headers := []string{
		"Bib", "Name", "School", "Finish", "Seed",
		"Varsity Points", "Seed Points",
		"Top 30% Points", "Dual Points", "Total Points",
	}

	if err := w.Write(headers); err != nil {
		return err
	}

	for _, r := range racers {
		row := []string{
			r.Bib,
			r.Name,
			r.School,
			strconv.Itoa(r.Finish),
			strconv.Itoa(r.Seed),
			strconv.Itoa(r.Varsity),
			strconv.Itoa(r.SeedPoints),
			strconv.Itoa(r.Top30Points),
			strconv.Itoa(r.DualPoints),
			strconv.Itoa(r.Total),
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}

	return w.Error()
}

func writeTeamCSV(path string, racers []*Racer) error {
	totals := make(map[string]int)
	names := make(map[string]string)

	for _, r := range racers {
		key := normalize(r.School)
		totals[key] += r.Total
		names[key] = r.School
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{"School", "Total Points"}); err != nil {
		return err
	}

	var keys []string
	for key := range totals {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return totals[keys[i]] > totals[keys[j]]
	})

	for _, key := range keys {
		row := []string{
			names[key],
			strconv.Itoa(totals[key]),
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}

	return w.Error()
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go run_order.csv individual_results.csv team_results.csv")
		os.Exit(1)
	}

	fmt.Println("===================================")
	fmt.Println("       SKI RACE POINT CALCULATOR")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("Select race type:")
	fmt.Println("1. Dual Meet")
	fmt.Println("2. Invitational / Sections")
	fmt.Print("Enter 1 or 2: ")

	var choice string
	fmt.Scanln(&choice)

	var raceType string

	switch strings.TrimSpace(choice) {
	case "1":
		raceType = "dual"
	case "2":
		raceType = "invitational"
	default:
		fmt.Println("Invalid selection.")
		os.Exit(1)
	}

	fmt.Println("\nReading run order...")

	racers, err := readRunOrder(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Reading individual results...")

	if err := readIndividualResults(os.Args[2], racers); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Reading team results...")

	teams, err := readTeamResults(os.Args[3])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Validate that every school in the run order
	// participated according to the team results.
	participating := make(map[string]bool)
	for _, t := range teams {
		participating[normalize(t.School)] = true
	}

	for _, r := range racers {
		if !participating[normalize(r.School)] {
			fmt.Printf("Error: school %s is in run order but not in team results\n", r.School)
			os.Exit(1)
		}
	}

	fmt.Println("\nCalculating points...")

	if raceType == "dual" {
		calculateDual(racers, teams)
	} else {
		calculateInvitational(racers)
	}

	calculateTotals(racers, raceType)

	fmt.Println("Writing individual results...")

	if err := writeIndividualCSV("individual_scored.csv", racers, raceType); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Writing team totals...")

	if err := writeTeamCSV("team_totals.csv", racers); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("\n===================================")
	fmt.Println("          RACE COMPLETE")
	fmt.Println("===================================")
	fmt.Printf("Race type: %s\n", raceType)
	fmt.Printf("Racers processed: %d\n", len(racers))
	fmt.Printf("Teams entered: %d\n", len(teams))
	fmt.Println("\nFiles created:")
	fmt.Println("  individual_scored.csv")
	fmt.Println("  team_totals.csv")
}
