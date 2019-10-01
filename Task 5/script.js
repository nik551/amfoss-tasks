function getdata() {
    var login = document.getElementById("nikme").value;
    var query = `query get_details($login: String!) {
                                user(login: $login) {
                                    login
                                    name
                                    avatarUrl
                                    bio
                                    email
                                }
                            }`
    fetch('https://api.github.com/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': 'bearer fe96880a614106a31bb41a2f31f379a9f96e39b8',
        },
        body: JSON.stringify({
            query,
            variables: {
                login
            },
        })
    }).then(r => r.json()).then(data => {
        console.log(data);
        
        if (data.data.user == null) {
            document.getElementById("error").style.display = "block";
        } else {
            document.getElementById("content").style.display = "block";
            document.getElementById("Avatar").src = data.data.user.avatarUrl;
            document.getElementById("Identicon").src = "https://identicons.github.com/" + document.getElementById("nikme").value + ".png";
            document.getElementById("Name").innerHTML = data.data.user.name
            document.getElementById("Bio").innerHTML = data.data.user.bio
            document.getElementById("Mail").innerHTML = data.data.user.email;
        }
    });
}