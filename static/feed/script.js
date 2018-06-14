const authToken = getCookie('socialnet_token')

Array.from(document.querySelectorAll('.post a')).forEach(el => {
  el.href = `/user/${el.getAttribute('data-id')}`
})

Array.from(document.querySelectorAll('.like')).forEach(el => {
  el.setAttribute('value', el.getAttribute('data-liked') === 'true' ? 'Unlike' : 'Like')
})

Array.from(document.querySelectorAll('.post .like')).forEach(el => {
  el.addEventListener('click', async () => {
    try {
     await fetch('http://localhost:3001/like', {
       method: 'post',
       headers: new Headers({
        'Content-Type': 'Application/json',
        'Authorization': 'Bearer ' + authToken
       }),
       body: JSON.stringify({ 
         postID: el.getAttribute('data-id'),
         like: el.value.toLowerCase() === 'like'
       })
     }).then(console.log)
    } catch (e) {
      console.error(e)
    }

    window.location.reload()
  })
})

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}

