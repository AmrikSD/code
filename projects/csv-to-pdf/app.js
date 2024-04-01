function readSingleFile(e) {
    const file = e.target.files[0];
    if (!file) {
      return;
    }   
    const reader = new FileReader();
    reader.onload = function(e) {
      const contents = e.target.result;
      displayContents(contents);
    };
    reader.readAsText(file, 'TIS-620');
}
  

function displayContents(contents) {

    const chunkSize = 30

    let regex = /,\s*(?=(?:[^"]|"[^"]*")*$)/g;

    let matrix = contents.split(regex)
        .map(str => str.replaceAll("\"", ""))
        .reduce((resultArray, item, index) => {
            const chunkIndex = Math.floor(index / chunkSize)

            if (!resultArray[chunkIndex]) {
                resultArray[chunkIndex] = [] // start a new chunk
            }

            resultArray[chunkIndex].push(item)

            return resultArray
        }, [])
        .filter(arr => arr.length >= 19)
        .map(item => {
            return {
                Order: item[2],
                Company: item[4],
                Date: item[6],
                Fax: item[10],
                Amount: `${item[18]} ${item[19]}`,
                Weight: `${item[20]} ${item[21]}`,
                Item: item[22],
                At: `${item[23]} ${item[24]}`,
                Note: item[25],
                Cost: item[26]
            }
        })

    const tableData = matrix.map(value => {
        return (
          `<tr>
             <td>${value.Amount}</td>
             <td>${value.Weight}</td>
             <td>${value.Item}</td>
             <td>${value.At}</td>
             <td></td>
             <td>${value.Note}</td>
             <td>${value.Cost}</td>
          </tr>`
        );
      }).join('');

    const tableBody = document.querySelector("#tableBody");
    tableBody.innerHTML = tableData;
    
    const orderNumber = document.getElementById("order-number")
    orderNumber.innerHTML = matrix[0].Order

    const date = document.getElementById("date")
    date.innerHTML = matrix[0].Date

    const companyName = document.getElementById("company-name")
    companyName.innerHTML = matrix[0].Company

    const companyFax = document.getElementById("company-fax")
    companyFax.innerHTML = matrix[0].Fax

    const companyLogo = document.getElementById("logo")
    const companyNameThai = document.getElementById("company-name-thai")
    const companyNameEnglish = document.getElementById("company-name-english")

    if (matrix[0].Order[2] === 'S'){
        companyLogo.outerHTML = "<img class='logo' id='logo' src='logos/ASG.png'>"
        companyNameThai.innerText = "บริษัท อภิโชคสตีล กรุ๊ป จำกัด"
        companyNameEnglish.innerText = "APICHOK STEEL GROUP CO., LTD."
    }


    const content = document.getElementById("content")
    content.innerHTML += content.innerHTML

    document.title = matrix[0].Order

}

  
document.getElementById('input').addEventListener('change', readSingleFile, false);