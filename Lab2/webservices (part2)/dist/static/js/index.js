Vue.component('items-box', {
    props: ['item'],
    template: '<div> <br> <li>{{ item.Id }} </li> <li>{{ item.value }} </li></div>'
})

var getItemsApp = new Vue({
    el: '#getItemsApp',
    data: {
        items: {}
    },
    methods: {
        getItems: function () {
            fetch("http://localhost:1337/api/getItems", {
                method: "GET",
                mode: "cors",
                credentials: "same-origin",
                redirect: "follow",
            })
            .then((res) => {
                return res.json();
            })
            .then((resJson) => {
                this.items = resJson.Items;
                console.log(resJson);
            })
        }
    }
})

var addItemApp = new Vue({
    el: '#addItemApp',
    data: {
        itemId: 1,
        itemValue: null
    },
    methods: {
        postItem: function () {
            fetch("http://localhost:1337/api/addItem", {
                method: "POST",
                mode: "cors",
                credentials: "same-origin",
                redirect: "follow",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Id: parseInt(this.itemId),
                    value: this.itemValue
                })
            })
            .then((res) => {
                return res.json();
            })
            .then((resJson) => {
                console.log(resJson);
                this.itemId++;
                this.itemValue = "";
            })
        }
    }
})