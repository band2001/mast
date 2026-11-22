import { Document } from "./document"

export interface Team extends Document {
    name: string,
    aliases: string[]
}