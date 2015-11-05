if ($Args.Length -ne 5) {
    echo "Invalid argument."
    exit 250
}

$option = $Args[0]
$remotehost = $Args[1]
$username = $Args[2]
$passwd = $Args[3]
$execution= $Args[4]

$password = ConvertTo-SecureString $passwd -Asplaintext -Force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)
$session = New-PSSession -ComputerName $remotehost -Credential $cred

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
    exit 250
}

$result = $?
echo $stdout

if($result -eq $false){
    Remove-PSSession -Session $session
     exit 250
}
    
$exitcode = Invoke-Command -Session $session -ScriptBlock { $lastexitcode }
Remove-PSSession -Session $session

exit $exitcode