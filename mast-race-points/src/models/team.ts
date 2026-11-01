interface Team extends Document {
    _id: string
    type: string
    createdAt: string
    updatedAt: string

    name: string,
    aliases: string[]
}