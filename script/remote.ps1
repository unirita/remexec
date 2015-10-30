$commandFlag = $Args[0]
$remotehost = $Args[1]
$username = $Args[2]
$passwd = $Args[3]
$execution= $Args[4]

if ($Args.Length -ne 5) {
    echo "Invalid argument."
    exit 1
}

$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)

$exitcode
$result
$stdout
$programexit

if ($commandFlag -eq "e"){
    $stdout = invoke-command -ComputerName $remotehost -Credential $cred -ScriptBlock{Invoke-Expression $args[0]; $exitcode = $lastexitcode } -argumentList $execution
}else{
    $stdout = invoke-command -ComputerName $remotehost -Credential $cred -FilePath $execution; $exitcode = $lastexitcode
}

$result = $?
echo $stdout

if ($exitcode -ne $nill){
    echo "ExitStatus: $exitcode"
    if($exitcode -eq 0){
        $programexit = 0
     }else{
        $programexit = 1
     }
}else {
    if($result -eq $true){
        echo "ExitStatus: 0"
        $programexit = 0
    }else{
        echo "ExitStatus: 1"
        $programexit = 1
    }
}

exit $programexit





