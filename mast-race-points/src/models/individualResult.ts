import { Document } from "./document"

export interface IndividualResult extends Document {
    raceId: string
    racerId: string
    finish: number
    seed: number
    finishInSeed: number
}