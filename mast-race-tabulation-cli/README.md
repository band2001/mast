
# Ski Race Points Calculator

A Go program that calculates individual and team points for high school ski races based on race results and varsity run order.

The program supports two race formats:

- **Dual Meets**
- **Invitationals / Sections**

It reads three CSV files, calculates points according to the race format, and generates two output CSV files containing individual skier points and team point totals.

---

## Table of Contents

1. [Requirements](#requirements)
2. [Installation](#installation)
3. [Running the Program](#running-the-program)
4. [Input CSV Files](#input-csv-files)
5. [Scoring Rules](#scoring-rules)
6. [Output CSV Files](#output-csv-files)
7. [Program Methods](#program-methods)
8. [Important Assumptions](#important-assumptions)
9. [Troubleshooting](#troubleshooting)
10. [Future Enhancements](#future-enhancements)

---

## Requirements

- Go 1.20 or newer recommended.
- Three CSV files:
  - Run Order
  - Individual Results
  - Team Results

No third-party Go packages are required. The program uses Go's standard library.

---

## Installation

1. Install Go from https://go.dev/dl/
2. Create a folder for the ski race calculator.
3. Save the program as `main.go`.
4. Save your three race CSV files in the same folder (or know their file paths).

Your folder might look like this:

```text
SkiRaceCalculator/
│
├── main.go
├── run_order.csv
├── individual_results.csv
└── team_results.csv
```

---

## Running the Program

Open a terminal or command prompt in the program directory.

Run:

```bash
go run main.go run_order.csv individual_results.csv team_results.csv
```

The program will display:

```text
===================================
       SKI RACE POINTS CALCULATOR
===================================

Select race type:
1. Dual Meet
2. Invitational / Sections

Enter 1 or 2:
```

Enter:

- `1` for a Dual Meet.
- `2` for an Invitational or Section race.

The program will read the files, calculate points, and generate the output CSV files.

### Compiling an Executable

Instead of running through Go each time, you can compile the program:

```bash
go build -o ski-race-calculator main.go
```

Then run:

**Windows:**

```bash
ski-race-calculator.exe run_order.csv individual_results.csv team_results.csv
```

**Mac/Linux:**

```bash
./ski-race-calculator run_order.csv individual_results.csv team_results.csv
```

---

## Input CSV Files

The program requires three CSV files.

### 1. Run Order CSV

This file determines which skiers were selected for varsity and their seed assignments.

Required columns:

| Column | Description |
|---|---|
| Bib | Racer's bib number |
| Name | Racer's name |
| School | Racer's school |

Example:

```csv
Bib,Name,School
101,John Smith,North High
102,Mike Jones,North High
103,Steve Brown,North High
201,Tom Wilson,South High
202,Chris Davis,South High
203,Matt Miller,South High
301,Sam Taylor,East High
302,Jack White,East High
303,Alex Green,East High
```

**Important:** The order of racers within each school determines their seed.

In the example above:

| Racer | School | Seed |
|---|---|---|
| John Smith | North High | 1 |
| Mike Jones | North High | 2 |
| Steve Brown | North High | 3 |
| Tom Wilson | South High | 1 |
| Chris Davis | South High | 2 |
| Matt Miller | South High | 3 |
| Sam Taylor | East High | 1 |
| Jack White | East High | 2 |
| Alex Green | East High | 3 |

The program assigns seeds automatically based on the order in which each school's racers appear.

### 2. Individual Results CSV

This file contains the actual race results.

Required columns:

| Column | Description |
|---|---|
| Bib | Racer's bib number |
| Name | Racer's name |
| School | Racer's school |
| Finish | Overall finishing position |

Example:

```csv
Bib,Name,School,Finish
101,John Smith,North High,2
102,Mike Jones,North High,5
103,Steve Brown,North High,8
201,Tom Wilson,South High,1
202,Chris Davis,South High,4
203,Matt Miller,South High,7
301,Sam Taylor,East High,3
302,Jack White,East High,6
303,Alex Green,East High,9
```

Finishing positions can be entered as:

- `1`
- `1st`
- `2`
- `2nd`

Non-numeric finishes such as DNF receive zero placement points.

Bib numbers must match the Run Order CSV.

### 3. Team Results CSV

This file contains the team standings for the race.

Required columns:

| Column | Description |
|---|---|
| School | School name |
| Finish | Team finishing position |

Example:

```csv
School,Finish
South High,1
North High,2
East High,3
```

The program uses this file for two purposes:

**Dual Meets:**
- Determines which schools participated.
- Determines team placement points for each school's skiers.

**Invitationals / Sections:**
- Determines which schools participated.
- Determines the number of participating teams, which is used to determine the number of seeds.

School names must match between the three CSV files.

---

## Scoring Rules

### Dual Meets

Each skier earns:

**1. Varsity Selection Point**

Every skier listed in the Run Order CSV receives 1 point.

**2. Team Placement Points**

Each skier receives points based on their school's team finish.

The number of points is determined by the number of participating teams.

For example, with three teams:

| Team Finish | Points Per Skier |
|---|---:|
| 1st | 3 |
| 2nd | 2 |
| 3rd | 1 |

The skier's total is:

```text
Total Points = Varsity Points + Dual Points
```

For example:

A North High skier whose team finishes second receives:

- 1 varsity point
- 2 dual points
- 3 total points

**Important:** The current implementation awards team placement points rather than calculating individual head-to-head finishes against each opposing school. See the Important Assumptions section.

### Invitationals / Sections

Each skier earns three types of points.

**1. Varsity Selection Points**

Every skier listed in the Run Order CSV receives 3 points.

**2. Seed Placement Points**

Seeds are determined by the order of racers within each school.

For example, if there are six schools with six racers each:

| Seed | Racers |
|---|---:|
| Seed 1 | 6 |
| Seed 2 | 6 |
| Seed 3 | 6 |
| Seed 4 | 6 |
| Seed 5 | 6 |
| Seed 6 | 6 |

All Seed 1 racers compete against each other for seed placement points, all Seed 2 racers compete against each other, etc.

Only the top half of each seed receives points.

Example: A seed with 18 racers.

| Seed Finish | Points |
|---|---:|
| 1st | 9 |
| 2nd | 8 |
| 3rd | 7 |
| 4th | 6 |
| 5th | 5 |
| 6th | 4 |
| 7th | 3 |
| 8th | 2 |
| 9th | 1 |
| 10th–18th | 0 |

The scoring formula is:

```text
Seed Points = (Seed Size / 2 rounded down) - Seed Rank + 1
```

Only skiers in the top half receive points.

**3. Overall Top 30% Points**

Skiers also earn points based on their overall finishing position.

The program calculates 30% of the total number of racers, rounding down.

Example: 100 racers.

| Overall Finish | Points |
|---|---:|
| 1st | 30 |
| 2nd | 29 |
| 3rd | 28 |
| 10th | 21 |
| 20th | 11 |
| 29th | 2 |
| 30th | 1 |
| 31st–100th | 0 |

For 35 racers, the top 10 receive points.

The scoring formula is:

```text
Top 30% Points = (Total Racers * 30 / 100) - Overall Rank + 1
```

Only skiers within the top 30% receive points.

**4. Invitational Total Points**

The final calculation is:

```text
Total Points =
    Varsity Points +
    Seed Points +
    Top 30% Points
```

---

## Output CSV Files

After successfully processing a race, the program generates two files.

### 1. individual_scored.csv

Contains the individual scoring breakdown for every racer.

Columns:

| Column | Description |
|---|---|
| Bib | Racer's bib number |
| Name | Racer's name |
| School | Racer's school |
| Finish | Overall finish |
| Seed | Assigned seed |
| Varsity Points | Points for varsity selection |
| Seed Points | Points earned within the seed |
| Top 30% Points | Points earned for overall placement |
| Dual Points | Dual meet team placement points |
| Total Points | Total points earned |

Example:

```csv
Bib,Name,School,Finish,Seed,Varsity Points,Seed Points,Top 30% Points,Dual Points,Total Points
101,John Smith,North High,2,1,3,2,2,0,7
102,Mike Jones,North High,5,2,3,0,0,0,3
```

Only the applicable scoring columns are populated for each race type.

### 2. team_totals.csv

Contains the total individual points earned by each school.

Columns:

| Column | Description |
|---|---|
| School | School name |
| Total Points | Combined points earned by all skiers from that school |

Example:

```csv
School,Total Points
North High,45
South High,39
East High,31
```

Teams are sorted from highest total points to lowest.

---

## Program Methods

The following describes the main functions in `main.go`.

### `normalize(s string) string`

Standardizes school names and CSV headers for comparisons.

It removes spaces, underscores, hyphens, periods, and number signs and converts text to lowercase.

This helps prevent minor formatting differences from causing matching problems.

### `column(headers []string, names ...string) int`

Finds the column index for a requested CSV field.

It allows certain variations in column names, such as:

- Bib
- Bib Number
- BibNumber

### `readCSV(path string) ([][]string, error)`

Opens and reads a CSV file.

Returns all rows from the CSV or an error if the file cannot be opened or read.

### `getField(row []string, index int) string`

Safely retrieves a value from a CSV row.

Returns an empty string if the requested column does not exist.

### `parseFinish(s string) int`

Converts a finishing position into an integer.

Examples:

- `1` becomes 1.
- `1st` becomes 1.
- `12th` becomes 12.
- `DNF` becomes 0.

### `readRunOrder(path string) ([]*Racer, error)`

Reads the Run Order CSV.

Responsibilities:

- Reads bib numbers, names, and schools.
- Checks for duplicate bib numbers.
- Assigns seeds based on each school's racer order.
- Creates the initial racer records.

### `readIndividualResults(path string, racers []*Racer) error`

Reads the Individual Results CSV.

Responsibilities:

- Matches individual results to racers using bib numbers.
- Validates school names.
- Reads finishing positions.
- Checks for duplicate or unexpected bib numbers.
- Assigns 1 varsity point to every racer in the run order.

### `readTeamResults(path string) ([]*Team, error)`

Reads the Team Results CSV.

Responsibilities:

- Reads school names and team finishing positions.
- Checks for duplicate schools.
- Creates team records.

### `assignRanks(racers []*Racer, getFinish func(*Racer) int, setRank func(*Racer, int))`

Sorts racers by finishing position and assigns ranks.

This method is used to calculate both seed placement and overall placement.

Ties are handled using the ranking method implemented in the function.

### `calculateInvitational(racers []*Racer)`

Calculates all invitational and section scoring.

Responsibilities:

1. Groups racers by seed.
2. Determines the number of racers in each seed.
3. Assigns seed finishing ranks.
4. Calculates seed placement points.
5. Assigns overall finishing ranks.
6. Calculates overall top 30% points.

### `calculateDual(racers []*Racer, teams []*Team)`

Calculates dual meet scoring.

Responsibilities:

1. Sorts participating schools by team finish.
2. Assigns team ranks.
3. Calculates team placement points.
4. Assigns those points to every skier on the corresponding team.

### `calculateTotals(racers []*Racer, raceType string)`

Calculates each racer's final race points.

For dual meets:

```text
Varsity Points + Dual Points
```

For invitationals:

```text
Varsity Points + Seed Points + Top 30% Points
```

### `writeIndividualCSV(path string, racers []*Racer, raceType string) error`

Creates the `individual_scored.csv` output file.

Writes each racer's information and scoring breakdown.

### `writeTeamCSV(path string, racers []*Racer) error`

Creates the `team_totals.csv` output file.

Adds together all individual racer points for each school and sorts teams by total points.

### `main()`

Controls the overall program.

Responsibilities:

1. Checks command-line arguments.
2. Asks the user to select a race type.
3. Reads all three input CSV files.
4. Validates participating schools.
5. Runs the appropriate scoring calculation.
6. Calculates individual totals.
7. Creates both output CSV files.
8. Displays a completion message.

---

## Important Assumptions

- Every racer in the Run Order CSV is considered a varsity selection.
- Every racer must have a matching bib number in the Individual Results CSV.
- School names must match across all three input files.
- Each school must appear only once in the Team Results CSV.
- The first racer listed for each school is Seed 1, the second is Seed 2, etc.
- Every participating school is represented in the Team Results CSV.
- The number of teams in the Team Results CSV determines the number of seeds available.
- DNF or other non-numeric finishes receive zero placement points.
- The current tie-ranking implementation uses dense ranking.
- Team totals are calculated by adding the points earned by every skier from that school.

### Important Scoring Considerations

The current implementation has two details worth reviewing before using it for official season standings:

1. **Dual scoring:** The original scoring description specifies individual placement points calculated separately against each opposing school. The current `calculateDual()` function instead awards points based on team finish. These are different scoring systems, and the dual scoring function should be updated if individual head-to-head scoring is intended.

2. **Invitational seed sizes:** Seed sizes are calculated from the actual number of racers in each seed. Each seed's actual size determines its scoring cutoff. This supports uneven team sizes.

---

## Troubleshooting

### Error: Missing CSV columns

Make sure the CSV headers contain the required fields.

The program recognizes variations such as:

- Bib / Bib Number
- Name / Racer Name
- School / Team
- Finish / Place / Position

### Error: Duplicate bib

A bib number appears more than once in the Run Order or Individual Results CSV.

Each racer must have a unique bib number.

### Error: School mismatch

A school name differs between the Run Order, Individual Results, and Team Results files.

Check spelling and remove unintended differences.

### Error: School in run order but not in team results

Every school represented in the Run Order CSV must also appear in the Team Results CSV.

### Error: CSV file cannot be opened

Verify that:

- The filename is correct.
- The file is in the specified directory.
- The file is not being used in a way that prevents access.
- The current user has permission to read the file.

---

## Future Enhancements

Potential future improvements include:

- Season-long point tracking across multiple races.
- Automatic accumulation of individual and team season standings.
- Support for multiple race dates.
- A graphical user interface.
- Automatic detection of race type.
- More detailed tie-handling options.
- Separate reporting for varsity selection points and placement points.
- Individual head-to-head dual meet scoring.
- Support for different team sizes and missing racers.
- An audit report showing exactly how every skier earned their points.

---

## License

This program is intended for ski team scoring and season-point tracking. Modify and distribute it as needed for your team's scoring requirements.