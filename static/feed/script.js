const authToken = getCookie('socialnet_token')

// Link user's name and image to their profile.
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

addPostListener(document.querySelector('#tweet-box'))

function addLikeListener(element, postID) {
  element.addEventListener('click', async () => {
    try {
      await fetch(serverAddress + '/like', {
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

function addCommentListener(buttonEl, inputEl, postID) {
  console.log('inputeEl: ', inputEl)
  buttonEl.addEventListener('click', async () => {
  console.log('fired comment listener', inputEl.value)
    try {
     await fetch(serverAddress + '/comment', {
       method: 'post',
       headers: new Headers({
        'Content-Type': 'Application/json',
        'Authorization': 'Bearer ' + authToken
       }),
       body: JSON.stringify({ 
         postID: postID,
         text: inputEl.value
       })
     }).then(console.log)
    } catch (e) {
      console.error(e)
    }

    window.location.reload()
  })
}

function addPostListener(newPostEl) {
  newPostEl.querySelector('button').addEventListener('click', async () => {
    const title = newPostEl.querySelector('.title').value 
    const image = newPostEl.querySelector('.pic').files[0]

    const formData = new FormData()
    formData.append('image', image)
    formData.append('title', title)

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
