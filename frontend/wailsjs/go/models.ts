export namespace main {
	
	export class Record {
	    name: string;
	    dep: string;
	    id: string;
	    bankid: string;
	    january: number;
	    february: number;
	    march: number;
	    april: number;
	    may: number;
	    june: number;
	    july: number;
	    august: number;
	    september: number;
	    october: number;
	    november: number;
	    december: number;
	    average: number;
	    total: number;
	    bonus: number;
	
	    static createFrom(source: any = {}) {
	        return new Record(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.dep = source["dep"];
	        this.id = source["id"];
	        this.bankid = source["bankid"];
	        this.january = source["january"];
	        this.february = source["february"];
	        this.march = source["march"];
	        this.april = source["april"];
	        this.may = source["may"];
	        this.june = source["june"];
	        this.july = source["july"];
	        this.august = source["august"];
	        this.september = source["september"];
	        this.october = source["october"];
	        this.november = source["november"];
	        this.december = source["december"];
	        this.average = source["average"];
	        this.total = source["total"];
	        this.bonus = source["bonus"];
	    }
	}

}

