async function _call(b) {
      return fetch('https://pkg.go.dev/play/compile', {
      method: 'POST',
      body: JSON.stringify({ body: b, version: 2 }),
    })
}

async function main() {
	var b;
	var result;
	process.argv.forEach(function (val, index, array) {
  		if (index == 2) {
			b = val;
	  		console.log(b);
			return
  		}
	});

	let r = await _call(b);
	let js = await r.json();
	let events = js['Events'];
	let errors = js['Errors'];
	if (errors.length === 0) {
		for (let i in events) {
			let e = events[i]['Message'];
			console.log(e);
		}
	} else {
		console.log(js['Errors']);
	}
}

main()

