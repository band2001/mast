import { Racer, RACER_TYPE } from "../models/racer"
import Repository from "../repository/Repository"

class RacerService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }

    async createRacers(racerInfo: { name: string, teamId: string }[]): Promise<Racer[]> {
        const racers: Racer[] = racerInfo.map(racer => (
            {
                _id: crypto.randomUUID(),
                type: RACER_TYPE,
                deleted: false,
                name: racer.name,
                teamId: racer.teamId,
                createdAt: Date.now(),
                updatedAt: Date.now()
            }
        ))
        const createResult = await this.repository.insertDocuments<Racer>(racers)
        const ids = createResult.insertedIds

        return this.repository.getDocumentsByProperty<Racer>("_id", ids)
    }

    async updateRacers(racers: Racer[]): Promise<Racer[]> {
        await this.repository.updateDocuments<Racer>(racers)
        
        const ids = racers.map(racer => racer._id)
        return this.repository.getDocumentsByProperty<Racer>("_id", ids)
    }
}