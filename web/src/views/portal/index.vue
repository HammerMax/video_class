<template>
    <div>
        <el-menu
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect">
            <el-menu-item index="0">
                <el-image
                    :src="url"
                    fit="scale-down"></el-image>
            </el-menu-item>
            <el-menu-item index="1">首页</el-menu-item>
            <el-menu-item index="2">搜索</el-menu-item>

        </el-menu>


    </div>
</template>

<script>
export default {
    data() {
        return {
            activeIndex: '1',
            activeIndex2: '1',
            url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        };
    },
    methods: {
        handleSelect(key, keyPath) {
            console.log(key, keyPath);
        }
    }
}
</script>
