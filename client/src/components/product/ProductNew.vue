<template>
    <div>
        <h3>New Product</h3>
        <p v-if="message" class="callout success">{{ message }}</p>
        <p v-if="error" class="callout alert">{{ error }}</p>
        <p>Name: <input type="text" v-model="item.Name"></p>
        <p>Description: <input type="text" v-model="item.Description"></p>
        <p>Price: <input type="number" v-model="item.Price"></p>
        <p><button class="button success" @click="saveProduct()">Save</button></p>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import ProductItem from '../product/ProductItem'
  export default {
      data() {
          return {
            item: {},
            message: "",
            error: ""
          }
      },
      methods: {
          saveProduct () {
              axios.post('/api/products', {
                name: this.item.Name,
                description: this.item.Description,
                price: parseInt(this.item.Price)
              })
              .then((res) => {
                  if(res.data.status){
                    this.item = res.data.data;
                    this.message = "Product saved successfully" 
                  }else{
                    this.error = "Something went wrong"  
                  }
              })
              .catch(error => console.log(error))
          },
      }
  }
</script>