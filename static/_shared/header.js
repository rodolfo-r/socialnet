const username = parseJWT(getCookie('socialnet_token')).usn
document.querySelector('#header-profile img').setAttribute('src', `http://localhost:3001/files/user/${username}/profile.jpg`)
document.querySelector('#header-profile a').setAttribute('href', `http://localhost:3000/user/${username}`)

document.querySelector('#logout').addEventListener('click', logout)

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}

function parseJWT(token) {
  const base64Url = token.split('.')[1]
  const base64 = base64Url.replace('-', '+').replace('_', '/')
  return JSON.parse(window.atob(base64))
}

function logout() {
 document.cookie = ''
  window.location = '/'
}
