// Your code goes here
//init buttons
var addButton = document.getElementsByClassName("add")
addButton[0].type = 'button' //disables refresh since value clears for reload
var submitButton = document.getElementsByTagName("button")[1]
submitButton.type = 'button'
var age_IP = document.getElementsByName("age")[0];
age_IP.type ='number' //make sure it is a number instead of validating
//list to store objects
candidateList = []

addButton[0].onclick = function(e) {
age_IP = document.getElementsByName("age")[0].value;
if(age_IP && (parseInt(age_IP,10) > 10)) 
	{
		var relationship_IP = document.getElementsByName('rel')[0].value;
		if(!relationship_IP){
		alert("select a relationship");
		return;
		}

	var smoker_IP = document.getElementsByName('smoker')[0].checked;
	var newCandidate =new Object();
	newCandidate.age = age_IP;
	newCandidate.relationship = relationship_IP;
	newCandidate.smoker = smoker_IP;
	candidateList.push(newCandidate);
	constructSomething();
	}
	else 
	{
	alert("age is invalid");
	}
}//end of add_button

function constructSomething() {
	var clearSomething = document.getElementsByClassName('household')[0].innerHTML = '';
	var something = document.createElement('ul');
	for(var candidates in candidateList) {
		var deleteBtn = document.createElement('button')
		deleteBtn.innerHTML =  "Remove"
		deleteBtn.onclick = function(currItem) {candidateList.splice(currItem, 1); constructSomething()}
		var item = document.createElement("li");
		var innerValue = candidateList[candidates].age+" years old and "+candidateList[candidates].relationship+", Smokes = "+candidateList[candidates].smoker;
		item.appendChild(document.createTextNode(innerValue))
		item.appendChild(deleteBtn);
		something.appendChild(item);
	}
	document.getElementsByClassName('household')[0].appendChild(something);	
}

submitButton.onclick= function() {
var debugElement = document.getElementsByClassName('debug')[0];
debugElement.style.display='block'
debugElement.innerHTML = JSON.stringify(candidateList); //element is present but it doesnt render
}	