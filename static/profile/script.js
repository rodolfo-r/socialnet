document.querySelector('#new-post .submit').addEventListener('click', async () => {
 const title = document.querySelector('#new-post .title').value 
 const body = document.querySelector('#new-post .body').value 

 const authToken = getCookie('socialnet_token')
  try {
   await fetch('http://localhost:3001/api/submit-post', {
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

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}
