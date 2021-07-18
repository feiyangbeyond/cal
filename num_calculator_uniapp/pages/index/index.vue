<template>
	<view class="page-root">
		<view class="con">
			<form @submit="formSubmit" @reset="formReset">
				<view class="uni-form-item uni-column" v-for="n in num" :key=n>
					<view class="title">第{{n}}个数字</view>
					<input class="uni-input" type="number" :name="'number'+ n" :placeholder="'请输入第'+n+'个数字'" />
				</view>
				<view class="uni-btn-v">
					<view style="width: 33.3%; margin: 0; float: left;">
						<button type="default" :disabled="loading" @click="numPlus">添加数字</button>
					</view>
					<view style="width: 33.3%; margin: 0; float: left;">
						<button type="primary" :loading="loading" :disabled="loading" form-type="submit">开始计算</button>
					</view>
					<view style="width: 33.3%; margin: 0; float: right;">
						<button type="primary" :disabled="loading" form-type="reset" plain="true">重新输入</button>
					</view>
				</view>
			</form>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: '请输入你要计算的数字',
				num: 1,
				loading: false,
				number: []
			}
		},
		onLoad() {

		},
		methods: {
			formSubmit: function(e) {
				this.loading = true
				var formdata = e.detail.value
				var nums = Object.values(formdata)
				var flag = true
				nums.forEach((n, i) => {
					if (!isNaN(n) && n !== '') {
						this.number.push(parseInt(n))
					} else {
						console.log(i)
						console.log('第' + (i + 1) + '个为非法数字:' + n)
						uni.showToast({
							icon: null,
							title: '第' + (i + 1) + '个为非法数字:' + n
						})
						flag = false
					}
				})

				if (!flag) {
					this.reset()
					return
				}

				// console.log(this.number)

				// console.log(sort(this.number))


				uni.showModal({
					content: '计算成功!',
					showCancel: true,
					confirmText: '导出excel',
					success: (res) => {
						if (res.confirm) {
							this.reset()
						} else {
							this.reset()
						}
					}
				});
			},
			formReset: function(e) {
				this.num = 1
				console.log('清空数据')
			},
			numPlus() {
				this.num++
			},
			reset() {
				this.loading = false
				this.number = []
			}
		}
	}
</script>

<style>
	.page-root {
		background-color: #F8F8F8;
		height: 100%;
		padding-bottom: 1px;
	}

	.con {
		/* margin: 10px 20px 10px 20px; */
	}

	.title {
		padding: 5px 13px;
	}

	.uni-form-item__title {
		font-size: 16px;
		line-height: 24px;
	}

	.uni-input-wrapper {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		padding: 8px 13px;
		flex-direction: row;
		flex-wrap: nowrap;
		background-color: #FFFFFF;
	}

	.uni-input {
		height: 40px;
		line-height: 40px;
		font-size: 15px;
		padding: 0px;
		flex: 1;
		background-color: #FFFFFF;
	}
</style>
