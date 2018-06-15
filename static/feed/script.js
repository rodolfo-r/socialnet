const authToken = getCookie('socialnet_token')

Array.from(document.querySelectorAll('.post a')).forEach(el => {
  el.href = `/user/${el.getAttribute('data-id')}`
})

Array.from(document.querySelectorAll('.like')).forEach(el => {
  el.setAttribute('value', el.getAttribute('data-liked') === 'true' ? 'Unlike' : 'Like')
})

Array.from(document.querySelectorAll('.post')).forEach(el => {
  const postID = el.getAttribute('data-id')

  const likeButton = el.querySelector('.like[type=button]')
  addLikeListener(likeButton, postID)

  const commentButton = el.querySelector('.comment [type=button]')
  const commentText = el.querySelector('.comment [type=text]')
  addCommentListener(commentButton, commentText, postID)
})

function getCookie(name) {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=')
  return parts.length == 2 ? parts.pop().split(';').shift() : ''
}

function addLikeListener(element, postID) {
  element.addEventListener('click', async () => {
    try {
      await fetch('http://localhost:3001/like', {
        method: 'post',
        headers: new Headers({
        'Content-Type': 'Application/json',
        'Authorization': 'Bearer ' + authToken
        }),
        body: JSON.stringify({ 
          postID: postID,
          like: element.value.toLowerCase() === 'like'
        })
      }).then(console.log)
    } catch (e) {
      console.error(e)
    }

    window.location.reload()
  })
}

function addCommentListener(buttonEl, textEl, postID) {
  buttonEl.addEventListener('click', async () => {
    try {
     await fetch('http://localhost:3001/comment', {
       method: 'post',
       headers: new Headers({
        'Content-Type': 'Application/json',
        'Authorization': 'Bearer ' + authToken
       }),
       body: JSON.stringify({ 
         postID: postID,
         text: textEl.value
       })
     }).then(console.log)
    } catch (e) {
      console.error(e)
    }

    window.location.reload()
  })
}
