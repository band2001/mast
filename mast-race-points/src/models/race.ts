import { Document } from "./document"

export interface Race extends Document {
    name: string
    venue: string
    classification: string
    numberOfRacers: number
    numberOfTeams: number
    competitionDate: string
    gender: string
    level: string
}