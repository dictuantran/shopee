import Configuration from './configuration';

class ItemService {

    constructor() {
        this.config = new Configuration();
    }

    async retrieveTaxcode(){        
        return fetch(this.config.TAXCODE_COLLECTION_URL)
            .then(response => {
                if(!response.ok) {
                    this.handleResponseError(response);                    
                }                
                return response.json();
            })
            .then(json => {                
                return json.data;
            })
            .catch(error=> {
                this.handleError(error);
            })
    }
    
    async retrieveItems() {
        return fetch(this.config.ITEM_COLLECTION_URL)
            .then(response => {
                if(!response.ok) {
                    this.handleResponseError(response);                    
                }
                return response.json();
            })
            .then(json => {
                const items = [];
                const itemArray = json._embedded.collectionItems;
                for(var i=0; i < itemArray.length; i++){
                    itemArray[i]["link"] = itemArray[i]._links.self.href;
                    items.push(itemArray[i])
                }
            })
            .catch(error=> {
                this.handleError(error);
            })
    }

    async getItem(itemLink) {
        return fetch(itemLink)
            .then(response => {            
                if (!response.ok) {
                    this.handleResponseError(response)
                }
                return response.json();
            })
            .then(item => {
                item["link"] = item._links.self.href;
                return item;
            })
            .catch(error => {
                this.handleError(error);
            });
    }

    handleResponseError(response) {
        throw new Error("HTPP error, status = " + response.status);
    }

    handleError(error) {
        console.log(error.message);
    }    
}

export default ItemService;