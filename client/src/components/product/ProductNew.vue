<template>
    <div>
        <h3>New Product</h3>
        <p>Name: <input type="text" v-model="item.Name"></p>
        <p>Description: <input type="text" v-model="item.Description"></p>
        <p><button class="button success" @click="saveProduct()">Save</button></p>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import * as types from '../../store/types';
  import ProductItem from '../product/ProductItem'
  export default {
      data() {
          return {
            item: {},
          }
      },
      methods: {
          saveProduct () {
              axios.post('/api/products', {
                name: this.item.Name,
                description: this.item.Description
              })
              .then((res) => {
                  this.item = res.data.data;
              })
              .catch(error => console.log(error))
          },
      }
  }
</script>