<template>
    <div>
        <ul class="list-group">
            <router-link
                    tag="li"
                    to="/categories/1"
                    class="list-group-item"
                    style="cursor: pointer">Category 1</router-link>
            <router-link
                    tag="li"
                    to="/categories/2"
                    class="list-group-item"
                    style="cursor: pointer">Category 2</router-link>
            <router-link
                    tag="li"
                    to="/categories/3"
                    class="list-group-item"
                    style="cursor: pointer">Category 3</router-link>
        </ul>
    </div>
</template>

<script>
import axios from '../../axios-auth';
import * as types from '../../store/types';
export default {
    data() {
        return {
            link: {
                name: 'categoryEdit',
                params: {
                    id: this.$route.params.id
                },
                query: {
                    q: 100
                },
                hash: '#data'
            },
            id: this.$route.params.id,
            categories: null
        }
    },
    methods: {
        getCategories () {
            axios.get('/api/categories')
            .then((res) => {
                this.$router.push({ name: 'categoryList' });
            })
            .catch(error => console.log(error))
        }
    },
    mounted () {
        this.getCategories()
    },
    beforeRouteEnter(to, from, next) {
        if (true) {
            next();
        } else {
            next(false);
        }
    }
}
</script>