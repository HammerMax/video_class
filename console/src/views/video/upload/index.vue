<template>
  <div class="app-container">
    <el-upload
      class="upload-demo"
      ref="upload"
      :auto-upload="false"
      :data="videoInfo"
      drag
      :file-list="fileList"
      :on-success="uploadSuccess"
      action="http://localhost:9999/storage/transcode">
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
    </el-upload>
    <el-form style="margin-top: 20px" ref="form" :model="videoInfo" label-width="80px">
      <el-form-item label="标题">
        <el-input v-model="videoInfo.title"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="createVideo">立即创建</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { videoApplyVid, videoCreate } from '@/api/video'

export default {
  data() {
    return {
      videoInfo: {
        vid: '',
        title: ''
      },
      fileList: []
    }
  },
  methods: {
    createVideo() {
      videoApplyVid().then(response => {
        if (response.code !== 0) {
          this.$message.error('申请vid失败')
          return
        }
        this.videoInfo.vid = response.data.vid
        this.$refs.upload.submit()
      })
    },
    uploadSuccess(response) {
      if (response.code !== 0) {
        this.$message.error(response.message)
        return
      }

      videoCreate(this.videoInfo).then(response => {
        if (response.code !== 0) {
          this.$message.error(response.message)
          return
        }

        this.$message.info('上传成功')
      })
    }
  }
}
</script>
