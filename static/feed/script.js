Array.from(document.querySelectorAll('#post a')).forEach(el => {
  el.href = `/user/${el.getAttribute('data-id')}`
})
