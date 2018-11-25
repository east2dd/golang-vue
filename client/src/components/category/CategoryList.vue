<template>

    <md-list class="md-double-line md-dense">
        <md-list-item v-for="item in items" :key="item.ID">
            <div class="md-list-item-text">
            <span>
                <router-link
                    tag="li"
                    :to="'/categories/' + item.ID"
                    class="list-group-item"
                    style="cursor: pointer"><a href="javascript:void(0)">{{ item.Name }}</a></router-link>
            </span>
            <span>id: {{ item.ID }}</span>
            <span>Description: {{ item.Description }}</span>
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
        getCategories () {
            axios.get('/api/categories')
            .then((res) => {
                this.items = res.data.data;
            })
            .catch(error => console.log(error))
        }
    },
    mounted () {
        this.getCategories()
    }
}
</script>