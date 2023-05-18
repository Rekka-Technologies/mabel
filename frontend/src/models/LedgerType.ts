

export class LedgerType {
    id: number;
    name: string;
    color: string;

    /**
     * Instead of hardcoding transaction types, we can use the backend to
     * provide us with a list of transaction types.
     * 
     * @param id Id of the ledger type
     * @param name Name of the ledger type
     * @param color Color to display the ledger type
     */
    constructor(id: number, name: string, color: string) {
        this.id = id;
        this.name = name;
        this.color = color;
    }

    /**
     * Lets parse the json data from the backend into a LedgerType object.
     * 
     * @param json Data to be parsed
     * @returns Parsed LedgerType object
     */
    static fromJson(json: any): LedgerType {
        if (!json) { return new LedgerType(0, '', ''); }
        return new LedgerType(json.id, json.name, json.color);
    }

    /**
     * Lets parse the json data from the backend into an array of LedgerType objects.
     * 
     * @param jsonArray Data to be parsed
     * @returns Array of parsed LedgerType objects
     */
    static fromJsonArray(jsonArray: any[]): LedgerType[] {
        if (!jsonArray) { return []; }
        return jsonArray.map((json) => LedgerType.fromJson(json));
    }
}