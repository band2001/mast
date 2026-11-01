import { Database, pocketDb } from "@axfab/pocket-db";
import path from "path";
import os from "os";

const initalizeDb = (): Database => {
    const appDataDir = process.env.APPDATA 
        || (
            process.platform === 'darwin' ? 
                path.join(os.homedir(), 'Library', 'Application Support') : 
                path.join(os.homedir(), '.config'));

    const userAppFolder = path.join(appDataDir, "mast-race-points");

    const dbPath = path.join(userAppFolder, "data.pdb");
    return pocketDb(dbPath);
}

const db = initalizeDb()

export default db