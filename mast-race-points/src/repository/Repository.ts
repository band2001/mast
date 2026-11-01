import { Collection, InsertManyResult, UpdateResult } from "@axfab/pocket-db"
import db from "./initializeDb"

class Repository {
    private static instance: Repository
    private documents: Collection

    private constructor() {
        this.documents = db.collection("documents")
    }

    static getInstance() {
        if (!Repository.instance) {
            Repository.instance = new Repository()
        }
        return Repository.instance
    }

    async getDocumentsByProperty<T extends Document>(property: keyof T, values: string[]): Promise<T[]> {
        return (this.documents.find({ [property]: { $in: values } }).toArray()) as unknown as T[]
    }

    async insertDocuments<T extends Document>(documents: T[]): Promise<InsertManyResult> {
        const payload = documents as unknown as Record<string, unknown>[]
        return this.documents.insertMany(payload)
    }

    async updateDocuments<T extends Document>(documents: T[]): Promise<UpdateResult[]> {
        return await Promise.all(
            documents.map(document => {
                const { _id, ...fieldsToUpdate } = document;

                return this.documents.updateOne(
                     _id,
                    fieldsToUpdate
                );
            })
        );
    }
}

export default Repository