interface TeamResult extends Document {
    _id: string
    type: string
    createdAt: string
    updatedAt: string

    raceId: string
    finish: number
    teamId: string
}