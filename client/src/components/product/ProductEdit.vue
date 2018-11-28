<template>
    <div>
        <h3>Edit the Product</h3>
        <p v-if="message" class="callout success">{{ message }}</p>
        <p v-if="error" class="callout alert">{{ error }}</p>
        <p><img src="https://via.placeholder.com/300.png"/></p>
        <p>Name: <input type="text" v-model="item.Name"></p>
        <p>Price: <input type="number" v-model="item.Price"></p>
        <p>Description: <input type="text" v-model="item.Description"></p>
        <p><button class="button success" @click="saveProduct()">Save</button></p>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import ProductItem from '../product/ProductItem'
  export default {
      data() {
          return {
            id: this.$route.params.id,
            item: {},
          }
      },
      methods: {
          saveProduct () {
              axios.put('/api/products/' + this.id, {
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
          getProduct () {
              axios.get('/api/products/' + this.id)
              .then((res) => {
                  this.item = res.data.data;
              })
              .catch(error => console.log(error))
          }
      },
      mounted () {
          this.getProduct()
      },
      beforeRouteLeave(to, from, next) {
          if (this.confirmed) {
              next();
          } else {
              if (confirm('Are you sure?')) {
                  next();
              } else {
                  next(false);
              }
          }
      }
  }
</script>