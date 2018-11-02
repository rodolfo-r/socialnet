let followBtn = document.querySelector('#follow')
const authToken = getCookie('socialnet_token')
const postEl = document.querySelector('#new-post')
const profileUser = document.URL.split('/')[document.URL.split('/').length -1]
const loggedUser = parseJwt(getCookie('socialnet_token')).usn
const serverAddress = 'https://socialnet.rodolforg.com/api'

if (loggedUser !== profileUser) {
  followBtn.type = "button"
}

const isFollower = document.querySelector('#follow').getAttribute('data-id')
if (isFollower === 'true') {
  followBtn.value = "Unfollow"
} else {
  followBtn.value = "Follow"
}

addFollowListener(followBtn)
addPostListener(postEl)

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}

function parseJwt (token) {
  const base64Url = token.split('.')[1]
  const base64 = base64Url.replace('-', '+').replace('_', '/')
  return JSON.parse(window.atob(base64))
}

function addFollowListener(element) {
  element.addEventListener('click', async () => {
    const shouldFollow = element.value.toLowerCase() === "follow"

    try {
      await fetch(`${serverAddress}/${shouldFollow ? '' : 'un'}follow`, {
        method: 'post',
        headers: new Headers({
          'Content-Type': 'Application/json',
          'Authorization': 'Bearer ' + authToken
        }),
        body: JSON.stringify({ username: profileUser })
    })
    } catch (e) {
      alert(e)
      return
    }

    window.location.reload()
  })
}

function addPostListener(newPostEl) {
  newPostEl.querySelector('.submit').addEventListener('click', async () => {
    const title = newPostEl.querySelector('.title').value 
    const body = newPostEl.querySelector('.body').value 
    const image = newPostEl.querySelector('.pic').files[0]

    const formData = new FormData()
    formData.append('image', image)
    formData.append('title', title)
    formData.append('body', body)

    const xhr = new XMLHttpRequest()
    xhr.open('post', serverAddress + '/submit-post', true)
    xhr.setRequestHeader('Authorization', `Bearer ${authToken}`)
    xhr.addEventListener('load', () => {
      window.location.reload()
    })
    xhr.addEventListener('error', evt => {
      console.error(evt)
    })
    xhr.send(formData)
  })
}
