import { Document } from "./document"

export interface Racer extends Document {
    name: string
    teamId: string
}

export const RACER_TYPE = "RACER"