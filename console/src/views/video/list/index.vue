<template>
  <div class="app-container">
    <div>
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="审批人">
          <el-input v-model="searchForm.user" placeholder="审批人"></el-input>
        </el-form-item>
        <el-form-item label="活动区域">
          <el-select v-model="searchForm.region" placeholder="活动区域">
            <el-option label="区域一" value="shanghai"></el-option>
            <el-option label="区域二" value="beijing"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div>
      <el-table
        :data="videoListTable"
        border
        style="width: 100%">
        <el-table-column
          prop="vid"
          label="Vid"
          width="180">
        </el-table-column>
        <el-table-column
          prop="image"
          label="封面图"
          width="180">
        </el-table-column>
        <el-table-column
          prop="title"
          label="标题"
          width="180">
          <template slot-scope="scope">
            <router-link :to="{ name: 'videoDetail', query: { vid: scope.row.vid }}"><el-link type="primary">{{ scope.row.title }}</el-link></router-link>
          </template>
        </el-table-column>
        <el-table-column
          prop="update_time"
          label="更新时间"
          width="180">
        </el-table-column>
        <el-table-column
          prop="create_time"
          label="上传时间">
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { videoList } from '@/api/video'

export default {
  data() {
    return {
      searchForm: {
        user: '',
        region: ''
      },
      videoListTable: undefined
    }
  },
  methods: {
  },
  mounted() {
    videoList().then(response => {
      this.videoListTable = response.data
    })
  }
}
</script>
