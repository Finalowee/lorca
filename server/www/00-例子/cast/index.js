const app = new Vue({
	el: "#app",
	data: {
	  books: [
	    {cnt:1, id:1, date:"2018-11-23", name:"Unix编程艺术", price:119},
	    {cnt:1, id:2, date:"2018-01-03", name:"代码大全", price: 105},
	    {cnt:1, id:3, date:"2016-09-07", name:"深入理解计算机原理", price: 98},
	    {cnt:1, id:4, date:"2013-12-16", name:"现代操作系统", price: 87},
	  ]
	},
	methods: {
		incr(b) {
			b.cnt ++
		},
		decr(b) {
			if (b.cnt>0) {
	 			b.cnt --
			} else {
			}
		},
		remove(idx) {
			this.books.splice(idx, 1)
		}
	},
	computed: {
		totalPrice() {
			return this.books.reduce((pre, book) => pre + book.price * book.cnt, 0)
		}
	},
	filters: {
		showPrice(price) {
			return '￥' + price.toFixed(2)
		}
	}
})