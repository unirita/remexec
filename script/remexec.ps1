if ($Args.Length -ne 5) {
    echo "Invalid argument."
    exit 255
}

$option = $Args[0]
$remotehost = $Args[1]
$username = $Args[2]
$passwd = $Args[3]
$execution= $Args[4]

$session
if ($username -ne "" -And $pass -ne ""){
    $password = ConvertTo-SecureString $passwd -Asplaintext -Force
    $cred = New-Object -TypeName System.Management.Automation.PSCredential -ArgumentList $username, $password
    $session = New-PSSession -ComputerName $remotehost -Credential $cred
}else{
    $session = New-PSSession -ComputerName $remotehost
}

if ($session -eq $null){
    echo "No exist credential."
    exit 255
}

$stdout

if ($option -eq "-exe"){
    $stdout = Invoke-Command -Session $session -ScriptBlock{Invoke-Expression $args[0]} -ArgumentList $execution
}elseif($option -eq "-file"){
    $exelist = -split $execution
    $exefile = $exelist[0]
    $exearglist =  @()
    for ($i = 1; $i -lt $exelist.Length; $i++){
        $exearglist += $exelist[$i]
    }
    $stdout = Invoke-Command -Session $session -FilePath $exefile -ArgumentList $exearglist 
}else {
    echo "Unkown option $option"
    exit 255
}

$result = $?
echo $stdout

if($result -eq $false){
    Remove-PSSession -Session $session
     exit 255
}
    
$exitcode = Invoke-Command -Session $session -ScriptBlock { $lastexitcode }

Remove-PSSession -Session $session

exit $exitcode