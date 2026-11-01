import Repository from "../repository/Repository";

class RaceResultsService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }
}