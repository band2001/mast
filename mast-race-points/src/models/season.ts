import { Document } from "./document"

export interface Season extends Document {
    title: string
    racerIds: string[]
    raceIds: string[]
}