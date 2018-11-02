const serverAddress = 'https://socialnet.rodolforg.com/api'

document.querySelector('#change-profile .submit').addEventListener('click', async () => {
  const authToken = getCookie('socialnet_token')

  const image = document.querySelector('#change-profile .pic').files[0]
  const formData = new FormData()
  formData.append("image", image)

  const xhr = new XMLHttpRequest()
  xhr.open('post', serverAddress + '/profile-picture', true)
  xhr.setRequestHeader('Authorization', `Bearer ${authToken}`)
  xhr.addEventListener('load', () => {
    window.location.reload()
  })
  xhr.addEventListener('error', evt => {
    window.location.reload()
    console.error(evt)
  })
  xhr.send(formData)
})

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}
