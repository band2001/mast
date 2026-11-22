import { Collection, InsertManyResult, UpdateResult } from "@axfab/pocket-db"
import db from "./initializeDb"

import { Document } from "../models/document"

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
        const datedDocuments = documents.map(document => ({ ...document, createdAt: Date.now(), updatedAt: Date.now() }))
        const payload = datedDocuments as unknown as Record<string, unknown>[]
        return this.documents.insertMany(payload)
    }

    async updateDocuments<T extends Document>(documents: T[]): Promise<UpdateResult[]> {
        return await Promise.all(
            documents.map(document => {
                const datedDocument = { ...document, updatedAt: Date.now() }
                const { _id, ...fieldsToUpdate } = datedDocument;

                return this.documents.updateOne(
                     _id,
                    fieldsToUpdate
                );
            })
        );
    }
}

export default Repository