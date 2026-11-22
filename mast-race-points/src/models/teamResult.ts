import { Document } from "./document"

export interface TeamResult extends Document {
    raceId: string
    finish: number
    teamId: string
}