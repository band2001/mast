import Repository from "../repository/Repository"

class RacerService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }
}