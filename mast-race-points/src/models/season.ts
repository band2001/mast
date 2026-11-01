interface Season extends Document {
    _id: string
    type: string
    createdAt: string
    updatedAt: string

    title: string
    racerIds: string[]
    raceIds: string[]
}