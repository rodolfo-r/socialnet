let followBtn = document.querySelector('#follow')
const profileUser = document.URL.split('/')[document.URL.split('/').length -1]

const loggedUser = parseJwt(getCookie('socialnet_token')).usn
if (loggedUser !== profileUser) {
  document.querySelector('#follow').type = "button"
}

const isFollower = document.querySelector('#follow').getAttribute('data-id')
if (isFollower === 'true') {
  followBtn.value = "Unfollow"
} else {
  followBtn.value = "Follow"
}

document.querySelector('#new-post .submit').addEventListener('click', async () => {
 const title = document.querySelector('#new-post .title').value 
 const body = document.querySelector('#new-post .body').value 

 const authToken = getCookie('socialnet_token')
  try {
   await fetch('http://localhost:3001/submit-post', {
    method: 'post',
    headers: new Headers({
      'Content-Type': 'Application/json',
      'Authorization': 'Bearer ' + authToken
    }),
     body: JSON.stringify({ title: title, body: body })
   })
  } catch (e) {
    alert(e)
    return
  }

  window.location.reload()
})

followBtn.addEventListener('click', async () => {
  const shouldFollow = followBtn.value.toLowerCase() === "follow"

  const authToken = getCookie('socialnet_token')
  try {
    await fetch(`http://localhost:3001/${shouldFollow ? '' : 'un'}follow`, {
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
