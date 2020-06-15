import Configuration from './configuration';

class ItemService {

  constructor() {
    this.items = [
      {link:1, name:"test1", summary:"Summary Test 1", year:"2001", country:"us", price:"1000", description:"Desc 1"},
      {link:2, name:"test2", summary:"Summary Test 2", year:"2002", country:"uk", price:"2000", description:"Desc 2"},
      {link:3, name:"test3", summary:"Summary Test 3", year:"2003", country:"cz", price:"3000", description:"Desc 3"},
    ];
    this.config = new Configuration();
  }

  async retrieveItems() {
      return Promise.resolve(this.items);
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
            console.log(json.data);       
            return json.data;
        })
        .catch(error=> {
            this.handleError(error);
        })
    }

  async getItem(itemLink) {
    for(var i = 0; i < this.items.length; i++) {
      if ( this.items[i].link === itemLink) {
        return Promise.resolve(this.items[i]);
      }
    }
    return null;
  }

  async createItem(item) {
    console.log("ItemService.createItem():");
    console.log(item);
    return Promise.resolve(item);
  }

  async deleteItem(itemId) {
    console.log("ItemService.deleteItem():");
    console.log("item ID:" + itemId);
  }

  async updateItem(item) {
    console.log("ItemService.updateItem():");
    console.log(item);
  }

  handleResponseError(response) {
    throw new Error("HTPP error, status = " + response.status);
  }

  handleError(error) {
      console.log(error.message);
  }    

}

export default ItemService;
