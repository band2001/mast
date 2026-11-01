interface Race extends Document {
    _id: string
    type: string
    createdAt: string
    updatedAt: string

    name: string
    venue: string
    classification: string
    numberOfRacers: number
    numberOfTeams: number
    competitionDate: string
    gender: string
    level: string
}