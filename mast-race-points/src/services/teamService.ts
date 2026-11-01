import Repository from "../repository/Repository"

class TeamService {
    repository: Repository
    
    constructor() {
        this.repository = Repository.getInstance()
    }
}