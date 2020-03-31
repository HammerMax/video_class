import request from '@/utils/request'

export function videoList() {
  return request({
    url: '/video/list',
    method: 'get'
  })
}

export function videoDetail(vid) {
  return request({
    url: '/video/detail',
    method: 'get',
    params: { vid: vid }
  })
}

export function videoUpdate(vid, fields) {
  fields.vid = vid
  return request({
    url: '/video/update',
    method: 'get',
    params: fields
  })
}

export function videoApplyVid() {
  return request({
    url: '/video/apply_vid',
    method: 'get'
  })
}

export function videoCreate(video) {
  return request({
    url: '/video/create',
    method: 'post',
    data: video
  })
}
