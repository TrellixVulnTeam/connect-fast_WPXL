export namespace ipfs {
	
	export class FileStatus {
	    name: string;
	    cat: string;
	    size: number;
	    cumlativeSize: number;
	    childBlocks: number;
	    fileType: string;
	
	    static createFrom(source: any = {}) {
	        return new FileStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.cat = source["cat"];
	        this.size = source["size"];
	        this.cumlativeSize = source["cumlativeSize"];
	        this.childBlocks = source["childBlocks"];
	        this.fileType = source["fileType"];
	    }
	}

}

export namespace main {
	
	
	export class FileStatus {
	    name: string;
	    path: string;
	    size: number;
	    last_modified: string;
	    is_dir: boolean;
	    sys: any;
	
	    static createFrom(source: any = {}) {
	        return new FileStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.size = source["size"];
	        this.last_modified = source["last_modified"];
	        this.is_dir = source["is_dir"];
	        this.sys = source["sys"];
	    }
	}

}

