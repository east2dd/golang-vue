<template>
<div class="category">
    <div class="row">
        <div class="columns medium-4" v-for="item in items" :key="item.ID">
            <p><img src="https://via.placeholder.com/300.png" style="width: 300px;"/></p>
            <router-link
                tag="p"
                :to="'/categories/' + item.ID"
                class="list-group-item"
                style="cursor: pointer"><a href="javascript:void(0)">{{ item.Name }}</a></router-link>
            <p>{{ item.Description }}</p>
        </div>
    </div>
</div>
</template>

<script>
import axios from '../../axios-auth';
export default {
    data() {
        return {
            id: this.$route.params.id,
            items: null
        }
    },
    methods: {
        getCategories () {
            axios.get('/api/categories')
            .then((res) => {
                this.items = res.data.data;
            })
            .catch(error => window.console.log(error))
        }
    },
    mounted () {
        this.getCategories()
    }
}
</script>

<style scoped>
  .category {
      margin-bottom: 2em;
  }
</style>