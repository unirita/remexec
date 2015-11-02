if ($Args.Length -ne 5) {
    echo "Invalid argument."
    exit 250
}

$method = $Args[0]
$remotehost = $Args[1]
$username = $Args[2]
$passwd = $Args[3]
$execution= $Args[4]


$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)
$session = New-PSSession -ComputerName $remotehost -Credential $cred


$stdout
$result
$exitcode

if ($method -eq "e"){
    $stdout = invoke-command -Session $session -ScriptBlock{Invoke-Expression $args[0]} -argumentList $execution
}else{
    $exelist = -split $execution
    $exefile = $exelist[0]
    $exearglist =  @()
    for ($i = 1; $i -lt $exelist.Length; $i++){
        $exearglist += $exelist[$i]
    }
    $stdout = invoke-Command -Session $session -FilePath $exefile -ArgumentList $exearglist 
}

$result = $?
    
if($result -eq $false){
    Remove-PSSession -Session $session
     exit 250
}
    
$exitcode = Invoke-Command -Session $session -ScriptBlock { $lastExitCode }
Remove-PSSession -Session $session

echo $stdout

exit $exitcode