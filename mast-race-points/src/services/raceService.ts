import Repository from "../repository/Repository"

class RaceService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }
}