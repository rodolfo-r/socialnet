const forms = ['signup', 'login']

const addClickListener = form => {
  document.querySelector(`#${form} .submit`).addEventListener('click', async () => {
    let postData = {}
    postData.username = document.querySelector(`#${form} .username`).value
    postData.password = document.querySelector(`#${form} .password`).value
    if (form === 'signup') {
      postData.firstName = document.querySelector(`#${form} .first-name`).value
      postData.lastName = document.querySelector(`#${form} .last-name`).value
      postData.email = document.querySelector(`#${form} .email`).value
    }

    let token
    try {
      token = await fetch(`http://localhost:3001/api/${form}`, {
        method: 'post',
        headers: new Headers({
          'Content-Type': 'Application/json'
        }),
        body: JSON.stringify(postData)
      })
      .then(res => res.json())
      .then(json => json.token)
    } catch (err) {
      alert(err)
      return
    }

    console.log('token: ', token)
    document.cookie = `socialnet_token=${token};`
    window.location = `/user/${postData.username}`
  })
}
  
forms.forEach(addClickListener)
