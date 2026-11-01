import Repository from "../repository/Repository"

class SeasonService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }
}