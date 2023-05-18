

export class LedgerRecord {
    id: number;
    date: string;
    name: string;
    reference: string | null;
    amount: number;
    tags: string[] | null;

    /**
     * The constructor for the LedgerRecord class.
     * 
     * @param date Date in ISO format (YYYY-MM-DD)
     * @param name Name of the transaction
     * @param reference Reference id of the transaction
     * @param amount Credit or debit amount of the transaction
     * @param tags Tags associated with the transaction
     */
    constructor(id: number, date: string, name: string, reference: string | null, amount: number, tags: string[] | null) {
        this.id = id;
        this.date = date;
        this.name = name;
        this.reference = reference;
        this.amount = amount;
        this.tags = tags;
    }

    /**
     * Lets parse the json data from the backend into a LedgerRecord object.
     * 
     * @param json Data to be parsed
     * @returns Parsed LedgerRecord object
     */
    static fromJson(json: any): LedgerRecord {
        if (!json) { return new LedgerRecord(0, '', '', null, 0, null); }
        return new LedgerRecord(json.id, json.date, json.name, json.reference, json.amount, json.tags);
    }

    /**
     * Lets parse the json data from the backend into an array of LedgerRecord objects.
     * 
     * @param jsonArray Data to be parsed
     * @returns Array of parsed LedgerRecord objects
     */
    static fromJsonArray(jsonArray: any[]): LedgerRecord[] {
        if (!jsonArray) { return []; }
        return jsonArray.map((json) => LedgerRecord.fromJson(json));
    }
}