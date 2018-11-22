<template>
    <md-list class="md-double-line md-dense">
        <md-list-item v-for="item in items" :key="item.id">
            <div class="md-list-item-text">
            <span>
                <router-link
                    tag="li"
                    :to="'/categories/' + item.ID"
                    class="list-group-item"
                    style="cursor: pointer"><a href="javascript:void(0)">{{ item.name }}</a></router-link>
            </span>
            <span>id: {{ item.ID }}</span>
            </div>
        </md-list-item>
    </md-list>
</template>
<script>
import axios from '../../axios-auth';
import * as types from '../../store/types';
export default {
    data() {
        return {
            id: this.$route.params.id,
            items: null
        }
    },
    methods: {
        getProducts () {
            axios.get('/api/products')
            .then((res) => {
                this.items = res.data.data;
            })
            .catch(error => console.log(error))
        }
    },
    mounted () {
        this.getProducts()
    }
}
</script>