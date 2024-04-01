function readSingleFile(e) {
    var file = e.target.files[0];
    if (!file) {
      return;
    }   
    var reader = new FileReader();
    reader.onload = function(e) {
      var contents = e.target.result;
      console.log(contents)
      displayContents(contents);
    };
    reader.readAsText(file, 'TIS-620');
}
  

function displayContents(contents) {

    const chunkSize = 30

    let regex = /,\s*(?=(?:[^"]|"[^"]*")*$)/g;

    matrix = contents.split(regex)
                  .map(str => str.replaceAll("\"",""))
                  .reduce((resultArray, item, index) => { 
                    const chunkIndex = Math.floor(index/chunkSize)
                  
                    if(!resultArray[chunkIndex]) {
                      resultArray[chunkIndex] = [] // start a new chunk
                    }
                  
                    resultArray[chunkIndex].push(item)
                  
                    return resultArray
                  }, [])
                  .filter(arr => arr.length >= 19)
                  .map(item => {
                    return {
                      Amount: `${item[18]} ${item[19]}`,
                      Weight: `${item[20]} ${item[21]}`,
                      Item: item[22],
                      Discount: `${item[23]} ${item[24]}`,
                      Note: item[25],
                      Cost: item[26]
                    }
                  })

    console.table(matrix)

    const tableData = matrix.map(value => {
        return (
          `<tr>
             <td>${value.Amount}</td>
             <td>${value.Weight}</td>
             <td>${value.Item}</td>
             <td></td>
             <td>${value.Discount}</td>
             <td>${value.Note}</td>
             <td>${value.Cost}</td>
          </tr>`
        );
      }).join('');

      const tableBody = document.querySelector("#tableBody");
      tableBody.innerHTML = tableData;
}

  
document.getElementById('input').addEventListener('change', readSingleFile, false);