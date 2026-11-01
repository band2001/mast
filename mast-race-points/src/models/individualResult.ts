interface IndividualResult extends Document {
    _id: string
    type: string
    createdAt: string
    updatedAt: string

    raceId: string
    racerId: string
    finish: number
    seed: number
    finishInSeed: number
}