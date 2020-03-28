<template>
	<div>
		<el-row class="nav-row" justify="space-between" type="flex">
			<el-col :span="6" class="logo-row">
				<img
					alt="logo"
					class="nav-logo"
					src="https://qzonestyle.gtimg.cn/aoi/sola/20200328174042_zlUGzPXUFk.png"
				/>
				<span class="nav-text">Hammer</span>
			</el-col>

			<el-col :span="6">
				<el-menu
					:default-active="activeIndex"
					@select="handleSelect"
					active-text-color="#eb2f96"
					background-color="#545c64"
					mode="horizontal"
					text-color="#fff"
				>
					<el-menu-item
						:index="item.id"
						:key="item.id"
						class="nav-item"
						v-for="item in menu"
					>{{item.name}}</el-menu-item>
				</el-menu>
			</el-col>
			<el-col :span="6" class="search-row">
				<el-input class="nav-search" placeholder="请输入内容" prefix-icon="el-icon-search" v-model="search"></el-input>
			</el-col>
		</el-row>
		<el-row class="banner-row" type="flex">
			<el-col :span="24">
				<el-carousel :interval="5000" arrow="always" height="500px">
					<el-carousel-item :key="`banner-${index}`" v-for="(item, index) in banner">
						<img :src="item" alt="banner" class="banner-img" />
					</el-carousel-item>
				</el-carousel>
			</el-col>
		</el-row>
		<div class="content-container">
			<el-row
				:key="`content-row${row}`"
				class="content-row"
				justify="space-between"
				type="flex"
				v-for="row in Math.ceil(content.length / 3)"
			>
				<el-col
					:key="`content${index}`"
					:span="8"
					v-for="(current, index) in [(row-1)*3, (row-1)*3 + 1, (row-1)*3 + 2]"
				>
					<el-card class="content-card" shadow="hover">
						<div class="content-card-header" slot="header">
							<span>{{content[current].title}}</span>
							<el-button class="content-button" icon="el-icon-edit" type="primary"></el-button>
						</div>
						<div
							:key="`content-inner${row}${index}-${id}`"
							class="content-card-body"
							v-for="(con, id) in content[current].content"
						>{{con}}</div>
					</el-card>
				</el-col>
			</el-row>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			menu: [
				{ name: '首页', id: 'menu-1' },
				{ name: '模块1', id: 'menu-2' },
				{ name: '模块2', id: 'menu-3' }
			],
			banner: [
				'https://qzonestyle.gtimg.cn/aoi/sola/20200328181431_N7UlOG9tdW.jpg',
				'https://qzonestyle.gtimg.cn/aoi/sola/20200328183024_MjVYfZcX6S.jpg',
				'https://qzonestyle.gtimg.cn/aoi/sola/20200328182735_Sptzc4ad4j.jpg'
			],
			content: [
				{
					title: '好想出去玩',
					content: ['我想去看油菜花', '啥时候能出门玩', '真是太难了。。。']
				},
				{
					title: '不想天天在家呆着了',
					content: ['啥时候能出门玩', '真是太难了。。。']
				},
				{
					title: '想出去吃好吃的',
					content: ['火锅', '披萨。。。', '鱼。。。。。。']
				},
				{
					title: '公司饭真的不好吃',
					content: ['1。。。。。', '2。。。。。', '3。。。。。']
				},
				{
					title: '写不下去了。。。',
					content: ['卒。。。。', '啥时候能出门玩', '真是太难了。。。']
				},
				{
					title: '就这样吧',
					content: ['啥啥啥的', '？？？？？？', '…………………………']
				}
			],
			activeIndex: '',
			search: ''
		};
	},
	methods: {
		handleSelect(key, keyPath) {
			console.log(key, keyPath);
		},
		init() {
			this.activeIndex = this.menu[0].id;
		}
	},
	mounted() {
		this.init();
		// this.loadData();
	}
};
</script>
<style lang="scss">
.nav-row {
	background-color: #545c64;
}
.logo-row {
	padding: 0 24px;
	display: flex;
	flex-direction: row;
	align-items: center;
}
.nav-logo {
	display: flex;
	width: 36px;
	height: 36px;
	margin-right: 8px;
}
.nav-text {
	display: flex;
	color: #fff;
	font-weight: bold;
	font-size: 22px;
}
.nav-item {
	min-width: 100px;
	font-size: 16px;
	text-align: center;
}
.search-row {
	display: flex;
	flex-direction: row;
	align-items: center;
	padding: 0 30px;
	justify-content: flex-end;
}
.nav-search {
	display: flex;
	border-radius: 6px;
	border: 1px solid #eb2f96;
	width: 230px;
}
.banner-row {
	overflow: hidden;
}
.banner-img {
	width: 100%;
	height: 100%;
}
.content-button {
	float: right;
	padding: 3px 0;
}
.content-container {
	padding: 10px 8px;
	color: #eb2f96;
	background-color: #545c64;
}
.content-card {
	margin-bottom: 8px;
	min-height: 200px;
	margin: 8px;
	background-color: #262626;
	border: 1px solid transparent;
	color: #fff;
	.el-card__header {
		border-bottom: 1px solid #eb2f96;
	}
	.el-card__body {
		padding-top: 10px;
		padding-bottom: 10px;
	}
}
.content-card:hover {
	border: 1px solid #eb2f96;
	cursor: pointer;
}
.content-card-header {
	font-size: 14px;
	font-weight: bold;
}
.content-button {
	background-color: transparent;
	border: none;
	color: #eb2f96;
}
.content-card-body {
	font-size: 12px;
	line-height: 2;
}
</style>
